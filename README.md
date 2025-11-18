# plasmactl-release

A [Launchr](https://github.com/launchrctl/launchr) plugin for [Plasmactl](https://github.com/plasmash/plasmactl) that creates git tags with automatically generated changelogs.

## Overview

`plasmactl-release` automates the release process by creating semantic version tags and generating changelogs from git commit history. It uses [git-chglog](https://github.com/git-chglog/git-chglog) to analyze commits and create professional changelogs.

## Features

- **Automatic Changelog Generation**: Creates changelogs from conventional commit messages
- **Git Tagging**: Creates and pushes semantic version tags
- **Preview Mode**: Review changelog before creating the release
- **Custom Tags**: Specify custom version numbers
- **Secure Credentials**: Stores git credentials in keyring

## Usage

### Interactive Release

```bash
plasmactl release
```

This will:
1. Analyze commit history since the last tag
2. Generate a changelog
3. Prompt for version number
4. Create and push the git tag

### Preview Mode

Review the changelog without creating a tag:

```bash
plasmactl release --preview
```

### Custom Tag

Specify a version number:

```bash
plasmactl release --tag v2.1.0
```

## Configuration

### Store Git Credentials

Store credentials for pushing tags:

```bash
plasmactl keyring:set release_git_user
plasmactl keyring:set release_git_password
```

Or use GitHub Personal Access Token:

```bash
plasmactl keyring:set release_git_password
# Enter your GitHub PAT when prompted
```

## Workflow

### Standard Release Process

```bash
# 1. Make changes and commit following conventional commits
git commit -m "feat: add new authentication method"
git commit -m "fix: resolve memory leak in processor"

# 2. Preview the changelog
plasmactl release --preview

# 3. Create the release
plasmactl release

# 4. Publish artifacts for the release
plasmactl package
plasmactl publish
```

### Integration with Package/Publish

```bash
# Complete release and publish workflow
plasmactl release --tag v1.2.0
plasmactl package   # Creates archive with commit SHA
plasmactl publish   # Uploads to artifact repository
```

## Conventional Commits

The changelog generation works best with [Conventional Commits](https://www.conventionalcommits.org/):

```bash
# Features
git commit -m "feat: add user authentication"
git commit -m "feat(api): implement rate limiting"

# Bug Fixes
git commit -m "fix: resolve null pointer exception"
git commit -m "fix(database): correct migration script"

# Breaking Changes
git commit -m "feat!: redesign API authentication"
git commit -m "feat(auth)!: remove deprecated login method"

# Other Types
git commit -m "docs: update README installation steps"
git commit -m "chore: update dependencies"
git commit -m "refactor: simplify error handling"
git commit -m "test: add unit tests for validator"
```

## Changelog Format

Generated changelogs include:

- **Version Number**: Semantic version (v1.2.3)
- **Release Date**: ISO 8601 format
- **Grouped Changes**: By type (Features, Bug Fixes, etc.)
- **Commit References**: Links to commit SHAs
- **Breaking Changes**: Highlighted separately
- **Contributors**: Automatically credited

Example output:

```markdown
## [v1.2.0] - 2025-01-15

### Features
- Add user authentication system
- Implement rate limiting for API endpoints

### Bug Fixes
- Resolve memory leak in data processor
- Correct database migration script

### Breaking Changes
- Redesign API authentication (requires client updates)
```

## Options

- `--tag <version>`: Specify custom version tag (e.g., `v1.0.0`)
- `--preview`: Preview changelog without creating tag
- `--user <username>`: Git username (or use keyring)
- `--password <password>`: Git password/token (or use keyring)

## Git Author Configuration

The release commits are authored as:

```
Author: ${USER}
Email: noreply@plasma.sh
```

Where `${USER}` is your system username.

## Best Practices

1. **Follow Conventional Commits**: Use semantic commit messages
2. **Preview First**: Always run with `--preview` before creating tags
3. **Semantic Versioning**: Follow [SemVer](https://semver.org/) (MAJOR.MINOR.PATCH)
   - MAJOR: Breaking changes
   - MINOR: New features (backward compatible)
   - PATCH: Bug fixes
4. **Protect Main Branch**: Ensure only authorized users can push tags
5. **Automate**: Integrate into CI/CD for automated releases

## Error Handling

### "No commits found"
Ensure you have commits since the last tag:
```bash
git log --oneline
```

### "Authentication failed"
Store valid git credentials:
```bash
plasmactl keyring:set release_git_user
plasmactl keyring:set release_git_password
```

### "Tag already exists"
Choose a different version number or delete the existing tag:
```bash
git tag -d v1.0.0
git push origin :refs/tags/v1.0.0
```

## Integration

### With GitHub Releases

After creating a tag, create a GitHub Release:

```bash
plasmactl release --tag v1.0.0
gh release create v1.0.0 --generate-notes
```

### With CI/CD

```yaml
# .github/workflows/release.yml
release:
  steps:
    - name: Create Release
      run: plasmactl release --tag ${{ github.ref_name }}
    - name: Build and Publish
      run: |
        plasmactl package
        plasmactl publish
```

## Documentation

- [Plasmactl](https://github.com/plasmash/plasmactl) - Main CLI tool
- [Plasma Platform](https://plasma.sh) - Platform documentation
- [Conventional Commits](https://www.conventionalcommits.org/) - Commit message format
- [Semantic Versioning](https://semver.org/) - Versioning scheme

## License

Apache License 2.0
