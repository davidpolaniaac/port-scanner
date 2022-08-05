# port-scanner
golang port scanner

Clone the repository from [GitHub repo](https://github.com/davidpolaniaac/port-scanner.git). 

## Setup Instructions

*   **Step 1: Change Directory**

Change the current working directory to the `port-scanner` directory.

For example, in the Windows command prompt, run the following command:

```
cd C:\Users\[username]\port-scanner

```
*   **Step 2: Run Scanner**

Run the scanner from go file:


```
USAGE
  $ go run main.go

OPTIONS
  --host=scanme.nmap.org 
  --start=22 
  --end=80

EXAMPLES
  $ go run main.go
  $ go run main.go --host=scanme.nmap.org --start=22 --end=25
  $ go run main.go --host=scanme.nmap.org
```
*   **Step 3: Observe the identified ports**

```
$ go run main.go --host=scanme.nmap.org --start=22 --end=25
scanning ports.....
Port 25 is open
Port 22 is open
total scan time: 115.487638ms
$
````