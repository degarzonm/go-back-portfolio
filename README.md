# Go-back-portfolio

This is a complete REST API written in Go with the Gin framework and MongoDB database.

## Endpoints

- Unregistered users:

  - `/register` (Create, creates new soldier, receives name, id, and password as input)
  - `/login` (Retrieves the JWT token given Id and password)

- Soldier | Officer:

  - `/my-fort` (Read, info about the fort the soldier belongs to)
  - `/my-commander` (Read, info about the commander of the fort)
  - `/check-agenda-of-day` (Read, info about the agenda of the fort)

- Officer:

  - `/officer/modify-agenda` (Update the agenda of the fort this officer is commander)
  - `/officer/check-general-plan` (Read, info about the plan given by the general of the fort)
  - `/officer/my-troops` (Read, Info about the soldiers who belong to the officer fort)

- General:

  - `/general/create-plan` (Create or Update the general's plan)
  - `/general/create-fort` (Create a new fort)
  - `/general/edit-fort/[Fort]` (Update a fort given its ID)
  - `/general/my-forts` (Read, info about all the forts of this general)
  - `/general/set-fort-commander/[Officer]/[Fort]` (Update, set the commander officer of a fort the general owns)
  - `/general/transfer-fort/[General]` (Transfer the fort to another general, remove the entry from the inner list)
  - `/general/my-troops/[Fort]` (Read, info about the troops of all forts the general has)
  - `/general/lost-fort/[Fort]` (Delete fort and delete all the soldiers belonging to it)

- Recruiter:

  - `/recruiter/ascend/[Soldier|Officer]` (Update, change soldier role to officer, and officer to General)
  - `/recruiter/transfer/[Soldier|Officer]/[Fort]` (Update, set the new fort ID into the soldier)
  - `/recruiter/jubilate/[Soldier|Officer|General]` (Delete, deletes a soldier, officer or general)
  - `/recruiter/edit-soldier/[Soldier|Officer|General]` (Update, change basic info of the soldier)
  - `/recruiter/release-me/` (Update, turns this recruiter role to officer)
  - `/recruiter/new-recruiter/[Officer]` (Update, turns officer into recruiter)