package repository

//go:generate sh -c "rm -rf mocks && mkdir -p mocks"
//go:generate ../../bin/minimock -i NoteRepository -o ./mocks/ -s "_minimock.go"
