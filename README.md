# csvEmptyFieldsRemover
Remove rows that contain empty fields

## Output

Two files to be expected.

If the input file is `test.csv`
The output files would be in the same folder with names `test.csv-non-empty`, and `test.csv-empty`

## How to run it

```shell script
./csvEmptyFieldsRemover -input=FILE_NAME -index=INDEXES
```

- `-input=` accepts a CSV file name.
- `-index=` accepts indexes of fields to be checked; index starting form 0

e.g.:


```shell script
./csvEmptyFieldsRemover -input=test.csv -index="1,3,9"
```

- Input file is `test.csv`
- Indexes of fields to be checked are `1` and `3`.

```shell script
./csvEmptyFieldsRemover -input=test.csv -index="0"
```

- Input file is `test.csv`
- Index of a field to be checked is `0`.

```shell script
./csvEmptyFieldsRemover -input=test.csv
```

- Input file is `test.csv`
- All fields are to be checked.