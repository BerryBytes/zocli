# Manifest file design (.yaml | .json)

**_NOTE: NOT ALL THE FIELDS NOTED HERE CAN BE EDITED OR UPDATED BY THE USER, AS THIS FILE IS FOR
DEMONSTRATION OF THE MANIFEST FILE, WHICH WORKS BOTH WAYS I.E. WHILE EDITING BY USER AND WHILE
PRINTING INFORMATION ABOUT ANYTHING FROM THE CLI._**

As, we will be creating the environment, projects and even application
using the manifest file, we must standardize the file so that we can
utilize the standard even while printing the output of various structs
on the console to the user.

### sample

```yaml
apiVersion: app.01cloud.io/v1
kind: Application
metadata:
    id: 1
    createdat: "2023-08-09T12:22:52.311842Z"
    name: My App
spec:
    active: true
    projectcode: "first"
    project:
        id: 1
        name: test-proj
    plugin:
        id: 1
        name: wordpress
    cluster:
        id: 1
        name: default
    chart:
    owner:
        id: 1
        first_name: admin
        last_name: admin
    git_repository_info:
    git_repository: ""
    git_repo_url:
    git_service: ""
    image_url: ""
    image_namespace: ""
    image_repo: ""
    image_service: ""
    service_type: 0
    active: true
    operator_package_name: ""
    status: "running"
```

As above we can see the information of the application, so let us briefly describe those.

### api version

Api version will contain the versioning of the api, to make the file consistent
and of we encounter any changes on the file or api on near future.

### kind

Kind will determine the nature of this file, the kind can be of any from
_"application"_, _"environment"_, or _"project"_. And possibly we will add
organization to the manifest kind on near future.

### metadata

The metadata will be responsible for holding any information about their creation,
structure, purpose, and dependencies.

Metadata can contain the following fields,

1. ID -> integer field representing the id, which will be used to perform any
   activity using the CLI,
2. Name -> string field which will contain the name, and in case of project the
   name can be utilized to perform activity,
3. CreatedAt -> datetime field for the creation date and time,
4. ProjectCode -> if any codes are present
5. active -> boolean field which will indicate the active status
6. status -> string field which will represent the status like, "running", "stopped", "pending"

With these information, the metadata will reflect every information of the
project/application/environment

### spec

Spec will contain all the specification related information which further elaborates
the application or environment or project.

Fields that can be on spec are as below,

1. **project** (available only on application manifest) ->
   it will contain the information related to the project where the application resides on.
   it can contain,
   _ID_ -> which represents the id of the project,
   _name_ -> which represents the name of the project
2. **application** (available only on the environment manifest) ->
   it contains the information related to the application where the environment resides on.
   it can contain,
   _ID_ -> which represents the id of the project,
   _name_ -> which represents the name of the project
3. **Variables**(project | environment) ->
   variables can be of type global scoped (project scoped) and environment scoped
   if the file is of type project manifest, then the variable field will contain the global
   variables and if the file is of type environment then, the environment scoped variables
   can be found in variable field
4. **organization**(project) ->
   orgranization field will help to determine the organization information of the project
   the fields which can reside in organization fields are,
   _ID_ -> which represents the id of ht organization,
   _name_ -> which represents the name of the organization
5. **owner**(project) ->
   owner field will represent the user id and name of the owner of the project. It will contain,
   _id_ -> user's id,
   _first_name_ -> user's firstname,
   _last_name_ -> user's lastname,
   _email_ -> user's email
6. **subscription**(project) ->
   subscription will contain the information related to the subscription which is used on the
   current project.
   Note that, fields other than id cannot be changed or edited, as this subscription field is
   solely made to give information regarding subscription to the user
   The fields which can reside inside of subscription are,
   _apps_: 2,
   _backups_: 5,
   _ci_build_: 50,
   _cores_: 1000,
   _cron_job_: 10,
   _data_transfer_: 102400,
   _disk_space_: 10240,
   _id_: 1,
   _memory_: 2048,
   _name_: "Basic Plan",
7. **plugin**(environment) ->
   Plugin field will contain the information ergarding the plugin which is currently used on the
   application

```yaml
apiVersion: app.01cloud.io/v1
kind: Project
metadata:
    id: 1 // output
    createdat: "2023-08-09T12:22:52.311842Z" // output
    name: My project //required
spec:
    active: true
    organization:
        id: 0
        name: default // output
    subscription:
        id: 1 //required
        name: Basic Subscription
    owner: // output
        id: 1
        first_name: admin
        last_name: admin
```
