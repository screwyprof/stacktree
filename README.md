# stacktree
A TDD Kata to deal with trees

## Description
Given a list of stack traces from standard in with the function
invocations running left to right, print to standard out an
indented diagram showing the count of invocations across all
stacks and the function name.

Original task is avalable [here](https://gist.github.com/jakekeeys/01b73d0569c29bee44b4232f9f046253)

Example input:

```
main, workloop, select
main, parse_args
main, workloop, parse_data, parse_entry
main, workloop, select
```

```
Example Output:
4 main
	3 workloop
		1 parse_data
			1 parse_entry
		2 select
	1 parse_args
```
