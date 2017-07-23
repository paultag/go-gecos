go-gecos
========

Read the GECOS field on a Linux machine. The GECOS field is the place in
`/etc/passwd` (or other `passwd`-like entry, such as LDAP and friends)
that stores the real-name (as exposed in `os/user.User.Name`), Room location,
Office Phone, Home Phone, and "other".

These entries can be set or changed using the `chfn(1)` command.
