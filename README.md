# StringMatchGame
Simple string match game in Golang. 

Let’s write a game!
The game engine first selects a code (string) of a certain length at random (and does not disclose this to the
player).
The player is allowed certain number of guesses to strategically guess the code.
After each guess, the game engine informs the player of the following:
1. Number of characters that are correct, but in the wrong place and
2. Number of characters that are correct and in the correct place.
Example : the game could start off with the computer generating the code : KMAFYM
The game engine starts by setting this code and informing the player that the code has 6 characters.
If the user guesses ABCDEF, then the engine informs the player that the he got 2 characters that are
present in the string but not in the right place. Then the user has to try again and so on ...
BONUS​ ( Think about this only if you finish the above question early ) :
As a followup to this question, write a computer based player, that also guesses the string alongside the
player and gets one chance after player's each guess.
Also show the player what the stats for the computer player are after every chance.


#Test case and sample outputs
#Without Prints

C:\Users\Manigandan D\Desktop\StringMatchGame>go run main.go
hyespd
Enter text to guss: heydps
heydps

You have gussed the correct letter but not in the correct position are:  4
You have gussed the correct letter in the correct position are:  2

C:\Users\Manigandan D\Desktop\StringMatchGame>

#----------------------------------------------

C:\Users\Manigandan D\Desktop\StringMatchGame>go run main.go
Enter text to guss: maniga
You have gussed the correct letter but not in the correct position are:  1
You have gussed the correct letter in the correct position are:  0

C:\Users\Manigandan D\Desktop\StringMatchGame>go run main.go
Enter text to guss: ajayss
You have gussed the correct letter but not in the correct position are:  0
You have gussed the correct letter in the correct position are:  0

C:\Users\Manigandan D\Desktop\StringMatchGame>go run main.go
Enter text to guss: karthi
You have gussed the correct letter but not in the correct position are:  0
You have gussed the correct letter in the correct position are:  1

C:\Users\Manigandan D\Desktop\StringMatchGame>

#----------------------------------------------


#With prints

C:\Users\Manigandan D\Desktop\StringMatchGame>go run main.go
qmoopt
Enter text to guss: aaaaaa
aaaaaa

You have gussed the correct letter in the correct position are:  0
You have gussed the correct letter but not in the correct position are:  0

C:\Users\Manigandan D\Desktop\StringMatchGame>go run main.go
bygasq
Enter text to guss: zzgzza
zzgzza

Pos:  3   a  -->  Pos:  5   a
97  This pressents
You have gussed the correct letter in the correct position are:  1
You have gussed the correct letter but not in the correct position are:  1

C:\Users\Manigandan D\Desktop\StringMatchGame>go run main.go
smlgbx
Enter text to guss: smlgbx
smlgbx

You have gussed the correct letter in the correct position are:  6
You have gussed the correct letter but not in the correct position are:  0

C:\Users\Manigandan D\Desktop\StringMatchGame>go run main.go
tirswj
Enter text to guss: itsrjw
itsrjw

Pos:  0   t  -->  Pos:  1   t
116  This pressents
Pos:  1   i  -->  Pos:  0   i
105  This pressents
Pos:  2   r  -->  Pos:  3   r
114  This pressents
Pos:  3   s  -->  Pos:  2   s
115  This pressents
Pos:  4   w  -->  Pos:  5   w
119  This pressents
Pos:  5   j  -->  Pos:  4   j
106  This pressents
You have gussed the correct letter in the correct position are:  0
You have gussed the correct letter but not in the correct position are:  6

C:\Users\Manigandan D\Desktop\StringMatchGame>
