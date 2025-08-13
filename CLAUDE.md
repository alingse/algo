# Claude Code Configuration

## Custom Commands

### /commit-code
Commits all current changes and pushes to remote repository.
- Runs `git add .`
- Creates commit with message "Update leetcode solutions"
- Runs `git push`

Usage: `/commit-code`