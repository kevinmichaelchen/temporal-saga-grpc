Output demo.gif

#Set Shell zsh
#Require glow

Set FontSize 14
Set Width 1200
Set Height 800
Set FontFamily Iosevka
#Set LineHeight 1.6
#Set LetterSpacing 6
Set Width 500
Set Height 300
Set Padding 30
Set TypingSpeed 0.05

# INSTRUCTION
Type "Spin up Jaeger and Temporalite"
Sleep 200ms
Ctrl+C

# EXECUTE
Type@100ms "docker-compose up -d"
Sleep 500ms
Enter

# SHOW OUTPUT FOR A BIT
Sleep 3s

# CLEAR SCREEN
Hide
Type "clear"
Enter
Show