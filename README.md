# Robots

This project was developed as part of a Draftea challenge.

## How to use:
- Clone the repository
        
        $ git clone https://github.com/NicolasBachs/robots.git

- Go to the root of the cloned repository:

        $ cd robots

- Execute:

        $ docker-compose up --build

- Now the application is running on `port 8080`. You can test the application with the following POST:

    ```
    curl --location --request POST 'http://localhost:8080/robots' \
    --header 'Content-Type: text/plain' \
    --data-raw '5 3
    1 1 E
    RFRFRFRFFFFFFF

    3 2 N
    FRRFLLFFRRFLL

    0 3 W
    LLFFFLFLFL'
    ```

## Extra:
In `./test/input/*` and `./test/output/*` we have some test cases to test that the application is working correctly. For example:

Input:
```
2 2
0 0 N
FFRFFRFFRFFF

0 0 N
FFRFFRFFRFFF
```

Output:
```
0 0 W LOST
0 0 W
```
