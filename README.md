# Chattronics

## How to run

There are 2 main ways to run chattronics: real and mocked.

Given that to run using the real GPT model you require an API key (which will
be used to charge you), it's best to use the mock model for testing purposes.

### Mocked
Just run `make run-mock` and follow the instructions in the terminal.

### Real GPT Model
First, setup you API key using a environment variable file. Create a file 
named `.env` in root and with the exact following format 
```.env
OPENAI_API_KEY=your-api-key-here
```

Then, just run
`make run-real model=gpt-4`
for GPT-4 or
`make run-real`
for GPT-3.5-Turbo.

To set the temperature just add a `temperatur=0.2` flag to the run make 
command. The default value is 0.1.

## Run results
Each execution will create a folder inside the `runs` folder. If this 
directory doesn't exist for you, it'll be created after executing.

Each run folder will have the following format
```
runs
└─ Mar08_14-57-30
    ├── diagram.svg
    ├── circuit files
    └── logs
        ├── console.logs
        ├── reviewer.json
        └── architect.json
```
Diagram, circuits and other outfiles are stored in the specific run folder,
while all logs are stored in `logs`. The `console.logs` file is basically
the stdout of the application, i.e., Debug, Info, Warn, Error and Fatal
logs will be sent here. The json logs are the interactions with GPT in 
structured form. The name of these json files are referring to the persona
of the model in question. 