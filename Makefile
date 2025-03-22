NAME=gossa
all:
	go build -o $(NAME)
	./$(NAME) -conf configs/config.yaml
