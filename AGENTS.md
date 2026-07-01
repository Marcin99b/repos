# AGENTS.md

## Project

`repos` is a small personal Cobra CLI (Go) for managing local git repositories from one folder. Commands live in `cmd/`, one file per command. Keep that in mind before writing docs — this is a tiny tool, not a product.

## Writing the README

- Describe what the tool actually does. Check `cmd/*.go` for the real commands, flags, and behavior before writing anything — never document a feature that isn't in the code.
- Match the README's size to the project's size. A three-command CLI gets a short README, not a landing page.
- Write like a person explaining the tool to a friend, not like marketing copy. No "seamless", "powerful", "effortless", "robust", no emoji headers, no bullet list of "Features" with perfectly parallel phrasing — that symmetry is what makes AI-written text feel off.
- One short paragraph of pitch is enough. Skip badges, tables of contents, and sections like "Why use this?" or "Roadmap" unless the project actually needs them.
- Show usage as real commands the user can copy-paste, each in its own fenced code block (` ```bash `), not inline text or a plain indented block.
- A closing line with a bit of personality is fine. A closing section is not.
- After writing, reread it and cut anything that only exists to make the tool sound bigger than it is.
