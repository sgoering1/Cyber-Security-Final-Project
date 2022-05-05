# Cyber-Security-Final-Project 

This project uses image stenography to encode Creative Commons Copyright informtion into images for publishing online. 

# How to use
The program takes in one command line argument which is the filepth to the users go/src file.
Looks like (C: username/go/src). Any other path will throw and error and exit the program. This is to keep things simple and not go to deep into filepath walking. 

- Images 
  Only JPEGs are supported for this currently. There are six folders in the the program file, each coresponds to the Creative Commomns Copyright and each has a text file with the copyright information.
  Images in the folders will be encoded with the copyright information in the text file and the new encoded image will be saved in the /img folder for ease of access in the HTML file. 
  
 
