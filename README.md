Go command-line interface application

Steps completed until current commit:

1. Created a "greetings" message, prompting the user to insert his/her name
2. Learned to use different OS functions such as Stdin, Getwd, ReadDir
3. Created getUser function - displays the logged user name, reading it from the system
   (same as "whoami" command in Linux/Unix)
4. Created getCurrentDir function - displays the path to the current directory where you invoke the program
   (works as "pwd" command)
5. Created listFiles function - displays all files in the current directory
   (works as "ls" command - plan is to expand the function to read the command as "ls -l")
6. Created displayPrompts function - displays all prompts 
7. Learned to work with promptui library, in order to create all prompts
8. Prompts are now mapped individually per each corresponding function
9. Prompt selection includes an index and result after selecting the desired option

What is not included yet:

- Option to go back to the main menu
- More commands
- Option to invoke the program from other paths (go.mod errors)
- Code revision and refactoring
