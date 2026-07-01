# repos

A small CLI for keeping git repositories in one place on your machine.

Instead of remembering where you cloned everything, you pick one folder as your "wd" (working directory), and every repo lives there.

## Install

```bash
go install github.com/Marcin99b/repos@latest
```

## Usage

Set the folder where your repos live:

```bash
repos wd ~/code
```

Check what it's currently set to:

```bash
repos wd
```

Clone a repo into it:

```bash
repos get https://github.com/user/repo
```

List what's already there:

```bash
repos ls
```

That's the whole tool. No config file to edit by hand, no extra flags to learn.