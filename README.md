I have created a web app using angularjs and GoLang (Server side code).
I have used Golang’s default web server.
To make webserver of golang operational please enter the following command in the corresponding folder
go run tag.go inputStruct.go outputStruct.go priorityQueue.go
Please wait until CMD/terminal shows “Done with storing data” message.
AngularJS:
Step 1: Please execute the following command “npm install” in Clarifai folder
	If it is still throwing errors. Please try executing the following commands
		npm install @angular/material –save
		npm install @angular/cdk –save
Step 2: please execute ng-serve .
Step 3: Go to localhost:4200 using google chrome.
Pease enter search word

Efficiency of retrieving data:
O(10) which is constant time. Since API have to display 10 most probable pictures. Stored the tag name and corresponding pictures in map of heap with key being the key name and heap contains the url of the image.
Storing data:
O(log10 * no of tags for each image) 
