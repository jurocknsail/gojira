# GoJira Release Notes

## Version 1.6.0 - 08/07/2024
* Added option to allow configuration by project id (gojira config projectId <value>).
* Added personal access token support for Jira authentication (gojira.exe tokenlogin)
* Upgrade of some dependencies.

## Version 1.5.0 - 02/10/2019
* Fixed allowed dod types, now gotten from the yaml config
* No more default assignee for DoD tasks

## Version 1.4.0 - 17/07/2019
* Now uses Go 1.12.

## Version 1.3.0 - 16/07/2019
* The subtasks.json file embedded in the program is removed. Now DoD is configurable in .gojira/dod.yaml for externalisation purposes.

## Version 1.2.0 - 31/05/2019
* Gojira is now a CLI Tool for Jira, based on the same framework as Gokub : cobra !

## Version 1.1.0 - 13/05/2019
* No need anymore for the subtasks.json file, it is embedded in the go program.

## Version 1.0.0 - 06/05/2019
* First release of Gojira, as an all-in-one tool to inject DoD in Jira.
