# Selenium Bot Go Version
## What is this?
This is a browser bot driven by a script engine inside. All the behavours such as click and input in the prepared JSON formatted script will be parsed and executed.
It is originally implemented in Python but now to implemented in Go for practice. Ref: https://github.com/tonyguo2010/BannerBot

## Why did I create it?
For the work needed, sometimes we will test an online web based product. The regression test will cause too much time manually. Then I asked myself, can I pass this boring and repeating work to a bot? Then the story happened.
Originally, I composed the script by myself. It is a painful task, far more than coding. It's really hard to say I enjoyed it until one day I learned that Selenium IDE could record the operations on the browser. It is not a big gap between my script and Selenium IDE. I used less than 1 hour to update my program to perfectly compatible to the new recorded script.

## How to record the script?
Well, for some secure reason, Google Store banned Selenium IDE. But Firefox is still supporing it. So you can install a Firefox Developer version with Selenium IDE plugin. 
Don't worry. The recorded script is fully JSON formatted. My program bot can be executed in a computer without Selenium IDE. 

So, Enjoy it!

## Install Dependencies
"go get github.com/tebeka/selenium" 
in sub folder libs, to install the selenium dependency


Visit https://googlechromelabs.github.io/chrome-for-testing/ to download the driver which is the same version with your Chrome. 

Since this small application is created on Chrome driver.

## How to use this tool?
Use Selenium IDE to record a script and save it as a JSON file. (Don't be confused by Selenium IDE's side file, it is a JSON file format)
Copy it into the folder and execute the following command:

go run .\main.go <script.json>

The program will launch the ChromeDriver and follow the recorded actions in the script.json file.
