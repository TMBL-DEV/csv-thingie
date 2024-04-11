# CSV-THINGIE

this is a simple barebones csv cli utility.

Currently there is not a lot in the tool but there are features that i want to implement

## Todo

- ❌ filtering by column value
- ❌ quickly spinning up sqlite database based on csv
- ❌ remove duplicate rows

## Currently working

- ✅ filtering row with regex

## installation

Compile it from scratch with
install modules: `go install`
build tool: `go build`

now you have the csv executable file in the project root.

You can point your path to it or you can move the file to `mv /usr/bin/csv`

tada now you can use the tool.

## Examples

command: `csv filter -p "~/Downloads/example.csv" -r someEmail@gmail\.com`

output: `Setting up header row from old -> new file
Csv has been filtered see output in:  /home/user/Downloads/example.filtered.csv`

---
