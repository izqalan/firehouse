# Firehouse

<img src="https://upload.wikimedia.org/wikipedia/commons/e/ea/MUTCD-MD_W11-8%281%29.svg" align="right"
     alt="Firehouse signal wikimedia" width="244" height="173">

Firehouse is a command-line interface (CLI) tool written in Go for managing and interacting with your Firebase project as an administrator. It provides a powerful and convenient way to perform various administrative tasks directly from your terminal.


## Usage

### Load service account
Firehouse uses the firehouse service -s command to load a service account. Here's the syntax:

```bash
firehouse service -s path/to/service-account.json
```
Replace `path/to/service-account.json` with the actual file path of your service account JSON key file. This file contains the credentials that Firehouse uses to authenticate with Firebase on your behalf.


### Verifying Loaded Service Account

Once you've loaded a service account, use the firehouse service get command to verify:
```bash
firehouse service get
```
This command should display the path to the currently loaded service account JSON file. If no service account is loaded, the output will be "Not Loaded".

Remember:

- Firehouse stores the loaded service account information in memory until you either load a different service account or exit the Firehouse CLI.
- Reloading the same service account won't affect the current session.

## Authentication

### Create

### Update

### Delete