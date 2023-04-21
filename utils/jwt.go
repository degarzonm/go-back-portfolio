package utils

import (
    "context"
    "fmt"
    "log"
    "os"
    "time"

    "github.com/degarzonm/go-back-portfolio/database"

    jwt "github.com/dgrijalva/jwt-go"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

// details is used to map the details of a soldier
type SignedDetails struct {
	Uid  string
	Name string
	City string
	jwt.StandardClaims
}

var soldierCollection *mongo.Collection = database.OpenCollection(database.DBClient, "soldier")

var SECRET_KEY string = os.Getenv("SECRET_KEY")

// GenerateAllTokens generates both teh detailed token and refresh token
func GenerateAllTokens(name string, uid string, city string) (signedToken string, signedRefreshToken string, err error) {
	claims := &SignedDetails{
		Name: name,
		Uid:  uid,
		City: city,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	refreshClaims := &SignedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(168)).Unix(),
		},
	}

	token, err1 := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))
	refreshToken, err2 := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(SECRET_KEY))

	if err1 != nil || err2 != nil{
		log.Panic(err)
		return
	}

	return token, refreshToken, err
}

//UpdateAllTokens renews the user tokens when they login
func UpdateAllTokens(signedToken string, signedRefreshToken string, userId string) {
    var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

    var updateObj primitive.D

    updateObj = append(updateObj, bson.E{"token", signedToken})
    updateObj = append(updateObj, bson.E{"refresh_token", signedRefreshToken})


    upsert := true
    filter := bson.M{"user_id": userId}
    opt := options.UpdateOptions{
        Upsert: &upsert,
    }

    _, err := soldierCollection.UpdateOne(
        ctx,
        filter,
        bson.D{
            {"$set", updateObj},
        },
        &opt,
    )
    defer cancel()

    if err != nil {
        log.Panic(err)
        return
    }

    return
}


//ValidateToken validates the jwt token
func ValidateToken(signedToken string) (claims *SignedDetails, msg string) {
    token, err := jwt.ParseWithClaims(
        signedToken,
        &SignedDetails{},
        func(token *jwt.Token) (interface{}, error) {
            return []byte(SECRET_KEY), nil
        },
    )

    if err != nil {
        msg = err.Error()
        return
    }

    claims, ok := token.Claims.(*SignedDetails)
    if !ok {
        msg = fmt.Sprintf("the token is invalid")
        msg = err.Error()
        return
    }

    if claims.ExpiresAt < time.Now().Local().Unix() {
        msg = fmt.Sprintf("token is expired")
        msg = err.Error()
        return
    }

    return claims, msg
}