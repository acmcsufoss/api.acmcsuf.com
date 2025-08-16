# api.acmcsuf.com contributing guide

Thanks for contributing!

## Getting Started

Take a look at [the readme](/README.md) for information on getting the project running
locally. Read [the developer docs](/developer-docs/api-testing.md) for some tips
on testing the api.

---

Don't be afraid to ask for help with anything on this project. DMs are open!

### Issues

- First check if the issue already exists
- Use appropriate tags to designate as a feature request, bugfix, feature
  addition, etc.
- Note if the change is for the cli or the api.
- Open discussions are encouraged - have an idea but not sure if it's worth
  implementing or how to do it? Open an issue!

### Pull Requests

- *Ensure ALL CI checks pass*. If one or more doesn't, fix the issue and `git push` to
  update the pr.
- *Avoid merge commits when possible*. Use `git rebase origin/main` to make sure
  your changes are on top of the most recent commit from main. If you
  already have an upstream branch, you will need to use  `git push
  --force-with-lease` instead of a regular push after doing this.
    - Occasionally rebasing will create merge conflicts out of thin air - in
      these cases it's fine to just merge.
- Don't forget to [link PR to issue](https://docs.github.com/en/issues/tracking-your-work-with-issues/using-issues/linking-a-pull-request-to-an-issue) if you are solving one.
- Please request a review from a current project maintainer or ping us on
  discord when your pr is ready to merge.
  - We may respond to your review with some suggestions and/or changes that
    should be made. Once you resolve these, update the pr and request another
    review.
- The Copilot PR review is a useful feature - it often catches easy to miss
  copy-paste errors, spelling/grammar mistakes, and other things. Be cautious
  though - it can make mistakes, and always validate the issues it finds.
