# Usage
```
go get github.com/kardianos/git-credential-static
git config credential.helper static
```

# Backend
Stores passwords in plain text file in home directory.
Useful for https git connections where your home directory is considered secure
and the https password is unique.
