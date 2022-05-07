# Cyber-Security-Final-Project 

This project uses image stenography to encode Creative Commons Copyright informtion into images for publishing online. 

# How to use
The program takes in one command line argument which is the filepth to the users go/src file.
Looks like (C: username/go/src). Any other path will throw and error and exit the program. This is to keep things simple and not go to deep into filepath walking. 

- Images 
  Only PNGs are supported for this currently. There are six folders in the the program file, each coresponds to the Creative Commomns Copyright and each has a text file with the copyright information.
  Images in the folders will be encoded with the copyright information in the text file and the new encoded image will be saved in the /img folder for ease of access in the HTML file. 
  Sample images are supplied but feel free to test out any .png files youd like. The larger the size the longer it takes, a 20MB file took ~45 seconds for refrence. Also the larger the image the bigger the message to encode can be. Those can be edited by simply editing the text files in the CC- folders.
  
 - CLI Flags
   -e C:/user/go/src: encodes images and boots website. MUST USE C:/user/go/src and program files must be in this directory
   -d /imagefilepath: decodes the image given and outputs CC to the console, can not be run when website is being served. 
   
 -Local host 
  Must run with a local webserver and not with local host, /main has to be the folder it runs out of. We used a chrome extension which ran it on 127.0.0.1:PORT which worked well (https://chrome.google.com/webstore/detail/web-server-for-chrome/ofhbbkphhbklhfoeikjpcbhemlocgigb). PORT is added on in the main function.
  
  
 
  
 
