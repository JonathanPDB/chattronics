# Chattronics

### Setup
First, setup you API key using a environment variable file. Create a file
named `.env` in root and with the exact following format
```.env
OPENAI_API_KEY=your-api-key-here
```

Then, just run
`make build`
`make run`

Default model is GPT-4-Turbo (gpt-4-1106-preview) and the default temperature is 0.8
They can be sepcified by adding the `model` and `temperature` flags in the makefile commands.

### Test setup

There are two type of tests: investigative and direct

The **investigative** scenario uses another GPT model to emulate the user, while the **direct** scenario sends all project
requirements in the beginning of the execution and, therefore, not allowing for any answers to the questions.

They can be run with
`make build-investigative`
`make run-investigative testCase=potentiometer iterations=1`

and


`make build-direct`
`make run-direct testCase=potentiometer iterations=1`

The test case is the porject to run the tests on.
Both can also take the `model` and `temperature` params.

