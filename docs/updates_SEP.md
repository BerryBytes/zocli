### LAST UPDATES

1. The module highlighted features that is well tested and works across the clouds
   Basic CRUD actions for Project, Environment, Organnization, Application
   Permission, and variables management on the project level
   Organization switch which will create a different context for the user to work on.
   Manifest file for creation of project, application

2. The limitations the module and its features have
   Single manifest file for creation of all organization, project, application and an environment

3. Known bugs
   Settings command on project doesn't have context value utilized -> This is because this portion was developed way before context appeared on our project.
   Variables on project doesn't have insert mode, and contains only update, delete and get

4. Todo and in progress stories
   Subscription / resources option for project, env and organization
   Permission, logger conf., variables, scheduler, domain conf., backups, jobs, CI/CD, startup scripts, addons, on environments
   Members, groups, registry, dns, cluster, plugins, chart repos, management on organization
   Activity logs, chart catalogs on organization
   Overview , webshell and logs of instances residing on evnrionment

5. Your wishlist for month of sep and oct to improve the modules
   Manifest file for creation of environment,
   Members, group, cluster, package management in organization
   Snapshot, restores and scheduled backups of environment
   Overview , webshell and logs of instances residing on evnrionment

6. Documentation plan and progress
   Documentation has been synced upto the application and the remaining portions are environment, context management (possibly a separate section too along with the sub-commans description on each commands), organization.

### CONCLUSION to work in priority

1. We will be creating a manifest file design for environment creation.
2. We will be working for cluster, package
3. Then we will move to webshells and logs
