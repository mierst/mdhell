# mdhell
A bot intended to aid in moderating the hell that is managing multiple discords. 

# Features

Users should be managable from any discord server from other users with a certain role (or set of roles?). Access to audit logs should only be available from servers marked as masters. 

## 1.0.0 Milestone

- [ ] logging
- [ ] warning
- [ ] muting
- [ ] banning
- [ ] auditing 

## Logging

All actions taken by moderators should create a log entry with:
- Log Id (Unique ID for each entry)
- Moderator's name (and id) 
- User's name (and id)
- Discord Server name (and id)
- Command
- Reason given
- Timestamp

```
!mdlog log [ log_id ]
!mdlog mod [ user ]
!mdlog user [ user ]
!mdlog server [ server ]
!mdlog command [ command ] # e.g. ban, warn, mute
!mdlog reason [ "reason" ] 
!mdlog timestamp [ start_timestamp ] [ end_timestamp ] # A single timestamp will find exact matches
```

### Features of Logging
- [ ] Logs should be queryable by any dimension
  - [ ] log
  - [ ] mod
  - [ ] user
  - [ ] server
  - [ ] command
  - [ ] reason
  - [ ] timestamp
- [ ] Reason query will be a case insensitive string pattern match
- [ ] Timestamps will be ISO-8601 compliant (YYYY-MM-DD HH:MM:SS) and in UTC
- [ ] Log query results will be displayed as a summary
- [ ] Exapnded Log results will report in registered channel
- [ ] A log channel should exist on the master server(s) that outputs when a log entry has been made

## Warning

```
!warn [ user | user_id ] reason
```

This should create a warn message to the user with reason and create a log entry.

### Features of Warning
- [ ] A message should be sent to the master logging channel(s) reporting on how many times this user has been warned in the last 30 days

## Muting

```
!mute [ user | user_id ] [ time_in_seconds ] reason
!unmute [ user | user_id ] reason
!mute report
```

This should mute the user in all associated discords by adding a muting role, create a warn message to the user with reason and create a log entry.

### Features of Muting
- [ ] Muting role to be applied can be registered globally
- [ ] Muting roles is removed automatically after time expires
- [ ] At muting time, a message should be sent to the master logging channel(s) reporting on how many times this user has been muted in the last 30 days
- [ ] At Unmuting time, a message should be sent to the master logging channel(s) reporting that a user has been umuted and give the original reason, or reason for explicit unmuting
- [ ] A muting report will give a list 

## Banning

```
!ban [ user | user_id ] reason
```

This should create a ban action on the associated server, and then propagate to all the other servers. 

## Auditing

```
!audit 
!audit fix
```

This should create a report that lists all joined servers and reports on mismatches for banned, muted and warned states.

```
Bans:
- X users are banned
- Y users are muted
- Z users have been warned in the last 30 days
- A actions been logged in the last 30 days
Server A (master)
- # users found who should be banned ([list of users])
- # users found who should be muted ([list of users])
Server B
- All users who should be banned are banned
- All users who should be muted are muted 
Server C
- All users who should be banned are banned
- There is no muting role registered for this server!
```

### Features of Auditing
- [ ] A nightly report will run at 0000 UTC and reported in the master logging channel

# Infrastructure

This service will require:
- [ ] Logging Database (pinned to the master)
- [ ] Master Database (pinned to the master)
  - [ ] Server table for each discord
  - [ ] State tables for each action
  - [ ] Role or Action association for each state
- [ ] Role or Action association for each state
- [ ] A global website to add/register bots on discord servers # Master server generates the shared key?


