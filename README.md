# cncfctl

like kubectl but for CNCF landscape :wink: 

## what is CNCFctl


## Installation

### Brew
```bash 
brew tap arjundandagi/brew
brew install cncfctl
```

### Compiling From source 
```bash
 git clone https://github.com/arjundandagi/cncfctl
 go build 
 ./cncfctl
```

## Usage:

You just have to set 2 environment variables
your Github Token 
The company username in github you want to search for
for example: `sumup` for sumup because it exists as https://github.com/sumup

```bash 
export GITHUB_TOKEN=githubtoken
export GITHUB_ORG=mycompany
```
then run it as 
```bash 
cncfctl
```
or in a single line

```bash 
GITHUB_TOKEN=githubtoken GITHUB_ORG=mycompany cncfctl 
```

### TOD0:
- [ ] Add testcases :sad:
- [ ] Add contributing guidelines
- [ ] Add initial commands like `cncfctl get` , `cncfctl explain` , `cncfctl list` on projects 
- [ ] Learn how to strcture the code neatly
- [ ] Make it cli tool using Cobra (or anyother library)
- [ ] Create complettion script for autocompleting. 


