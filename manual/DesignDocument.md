# Architecture Decision Record: Commands Design and Persistent Context Management in zocli

Title: Commands Design and Context Switch and Management

Status: Proposal

## Context

As of right now, we are managing and developing **zocli** commands as like below,

1. Get Projects

```bash
zocli project get [id/name]
```

This command retrieves the list of projects if no **id** or **name** is supplied.
And If any **id** or **name** is supplied as an argument to the command, then that specific
the project will be fetched.

2. Get Applications

```bash
zocli app list [projectID/projectName]
```

This command retrieves the list of applications of the specific project whose **name** or
**id** is supplied as the argument.

This is because the applications on our eco-system reside inside of projects and we cannot
fetch the list without the project specification. So, for fetching single application or
project we can use their own respective ID but while fetching the list we must supply
the upper hierarchy id as then only the list can be fetched.

## Problems detected

> With the above context, let us now briefly look at the problems that can arise.

Now this is from where the problem arose as if we must supply the upper hierarchy **id**
or **name** on every command, then the command may seem to be like,

```bash
zocli env list <:applicationID>
```

As, the **applicationID** should be supplied every time the env is fetched, it becomes
very tricky as, many other commands need upper hierarchy's ID to be supplied.

<!-- and even worse when the organization id is to be supplied like, -->
<!--  -->
<!-- ```bash -->
<!-- zocli env get <orgID/orgName> <projectID/projectName> <applicationID> -->
<!-- ``` -->

## Purposal

> Basics of Git and Kubectl will be expected for this proposal

As like in the git command, we manage the context using the branches, and modern shells like,
**zsh** or **fish** can also detect the working branch for the folder we currently reside in.

##### Takeaway

So like that, we can configure or save the current working organization, project, and even
application information to be used on every command which requires the information for this id.

Now coming to kubectl, we can manage our context using the config subcommand like,

```bash
kuebctl config get-contexts

kubectl config current-context
```

And we can set or change our context, using **set-context** sub-command, which is available in kubectl.
Likewise, we can also configure various other default configs for running kubectl.

##### Takeaway

So, like kubectl we can also add a sub-command for managing our context for working with our eco-system.
The configs can be the default organization, projects, applications, or even the environment. So, with this,
our command manageability can outperform, if any of our users can use the context configuration.

##### Small Problem that may arise

Imagine a situation, where the user does not know about the feature of context or if the user is not
willing to use the context configuration, then we can also consider the below points,

1.  **Add the user session on bash**
    That is, if the user is using zocli and performing tasks, then the project or organization
    or even an environment that is used recently can be added to the bash, using the below-given commands,
    With this, we can maintain the user's session as long as the bash is not exited.

    ```bash
        export ZO_DEFAULT_ORGANIZATION="abc"

        export ZO_DEFAULT_PROJECT="xyz"
    ```

2.  Or, By adding temporary context configuration on the **/tmp** folder
    With this, the user context can be maintained, as long as the system is not turned down.

## Rationale

1. **Easy switch of context and work**: We can create different profiles, or context for
   working on different scenarios, with which we can leverage the previously created configurations
   and start working right away.

2. **Easier to stay focused on task**: If we implement this feature on our cli, then it will
   be easier for our users to maintain their focus on a single task.

3. **Single Person different profiles**: Using this any number of users can work on any
   number of organizations, and can switch between them at ease.

4. **Less sub-commands or argument combination**: Using this our commands will turn less in
   size, which will be easy for our users to execute them, and memorize them.

## Implications

1. **Pre-Existing Architecture**: This feature must be added on pre-existing codebase
   and logic. There will be a significant amount of changes for the addition of this logic on every command
   and sub-commands ever created.

2. **Ease to use**: After this feature gets implemented, zocli overall will be easy to use.
   And so, will be easy to work around our eco-system of ZeroOne Cloud.

## Conclusion

By implementing this feature this early can help us get the pace too soon, and because
of this our zocli will scale, as there will be many other things that can be configured
using the config subcommand other than context.
