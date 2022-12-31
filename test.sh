start_time="$(date -u +%s.%N)"
echo "Starting tests..."

#test cases from the audit page and more

# fail cases
echo go run . --color red "banana" 
echo
go run . --color red "banana" 
echo go run . -color red "banana" 
echo
go run . -color red "banana" 
echo go run . -color=red "banana" 
echo
go run . -color=red "banana" 
# test cases
echo go run . --color=red "hello world"
echo
go run . --color=red "hello world"
echo go run . --color=green "1 + 1 = 2"
echo
go run . --color=green "1 + 1 = 2"
echo go run . --color=yellow "(%&) ??"
echo
go run . --color=yellow "(%&) ??"
echo go run . --color=blue "ascii" 2:
echo
go run . --color=blue "ascii" 2:
echo go run . --color=magenta "working hard" o i a
echo
go run . --color=magenta "working hard" o i a
echo go run . --color=orange "HeY GuYs" 4:
echo
go run . --color=orange "HeY GuYs" 4:
echo go run . --color=brightBlue 'RGB()' B 
echo
go run . --color=brightBlue 'RGB()' B 
echo go run . --color=backgroundBrightWhite "XsTcDfFh"
echo
go run . --color=backgroundBrightWhite "XsTcDfFh"
echo go run . --color=brightGreen "4s7 a9d 8k" 
echo
go run . --color=brightGreen "4s7 a9d 8k" 
echo go run . --color=brightYellow "!@$%^&*"
echo 
go run . --color=brightYellow "!@$%^&*"
echo go run . --color=cyan "B 9Xz jfK 7" 9 X z 7
echo
go run . --color=cyan "B 9Xz jfK 7" 9 X z 7
echo go run . --color=brightRed "Get shit done" G e t i d o n
echo
go run . --color=brightRed "Get shit done" G e t i d o n

end_time="$(date -u +%s.%N)"
echo "Tests finished!"
elapsed="$(awk "BEGIN { print $end_time - $start_time }")"
echo "Total of $elapsed seconds elapsed for processing"