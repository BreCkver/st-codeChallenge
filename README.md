# Title: st-codeChallenge
## Overview: 
This program is going to receive input file, in csv format, from a mounted directory. The file
will contain a list of debit and credit transactions on an account, so should process the file
and send summary information to a user in the form of an email

#### Use Cases
The app coverage the next use cases..
* The file need almost 3 values such as: Identifier, Date and amount
* The formar should be MM/dd 
* Credit transactions are indicated with a plus sign like +60.5. 
* Debit transactions are indicated by a minus sign
* File processing can be partial, but should be existing at least a successful transaction
* Fields in the form everyone is mandatories
* File shouldn't empty  
* Stored in database

#### Out of Scope
* Transaction date and its validation then can be existed transaction from future
* It's obvious that transaction was decrease/increase the balance account already, so there isn't validation about balance account
* Files different format to csv

## Project structure
The architecture in this moment is simple, so it has the next structure:
data                -> data access
handlers            -> handler to any request 
images              -> images to email template
internal            -> program core functions
models              -> program models 
routers             -> functions that process request from handler
templates           -> templates that render in web browser or client email
uploads             -> Directory to back up the files processed

## Libraries used directly
- Gorilla           -> Handler request
- Mongo-driver      -> Mongo access
- GoMail            -> Email Sender

NOTE: It's import say that app was pull in dockerhub, so you can find it in breckver/nbchallenge

## How to compile and run 
1. Make a new direcotory
2. Change to new directory
3. Decompress the solution
4. Change to nbchallege directory where the sln protect be
5. Execute the commands:
	dotnet build 
	dotnet run

NOTE:
The solution has one file testing already, so this file you can find it in Uploads its name is transaction-.csv