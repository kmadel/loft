---
title: "Command - loft wakeup space"
sidebar_label: loft wakeup space
---


Wakes up a space

## Synopsis


```
loft wakeup space SPACE_NAME [flags]
```

```
#######################################################
################### loft wakeup space #################
#######################################################
wakeup resumes a sleeping space

Example:
loft wakeup space myspace
loft wakeup space myspace --project myproject
#######################################################
```


## Flags

```
      --cluster string   The cluster to use
  -h, --help             help for space
  -p, --project string   The project to use
```


## Global & Inherited Flags

```
      --config string   The loft config to use (will be created if it does not exist) (default "$HOME/.loft/config.json")
      --debug           Prints the stack trace if an error occurs
      --silent          Run in silent mode and prevents any loft log output except panics & fatals
```

