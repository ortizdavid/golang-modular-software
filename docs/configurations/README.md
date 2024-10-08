Here's the updated `README.md` file for the Configurations Module with sections for Module Flags and Feature Flags included:

---

# Configurations Module

## Overview
The Configurations Module provides a centralized interface for managing key settings across the application. It enables administrators to easily configure essential parameters that influence the behavior, appearance, and communication capabilities of the system.

## Features

- **Basic Configuration:** Manage core settings that define the application's identity and user experience.
    - **App Name:** Set the full name of the application for branding purposes.
    - **Pagination:** Configure default pagination settings to control the number of items displayed per page.
    - **Style:** Customize the application’s appearance, including themes and layouts.
    - **App Acronym:** Define a short form or acronym for the application, used in headers and abbreviations.
  
- **Email Configuration:** Configure email communication settings to ensure reliable delivery of notifications and other correspondence.
    - **SMTP Sender:** Specify the default sender email address for outgoing emails.
    - **SMTP Port:** Set the port used for SMTP communications.
    - **Password:** Securely store and manage the SMTP server password.
  
- **Company Configuration:** Define and maintain company-specific settings that are reflected across the application.
    - **Name:** Set the official name of the company.
    - **Acronym:** Define the company’s acronym for use in documents and branding.
    - **Phone:** Specify the primary contact phone number for the company.
    - **Email:** Set the main contact email address for the company.

- **Module Flag:** Manage the activation status of different modules within the application.
    - **Modules:** Enable or disable specific modules to control their availability and functionality.
    - **Status:** Set to `Enabled` or `Disabled` to reflect the module’s operational state.

- **Feature Flag:** Manage the activation status of specific features within modules.
    - **Features:** Enable or disable individual features to control their availability and functionality within modules.
    - **Status:** Set to `Enabled` or `Disabled` to reflect the feature’s operational state.

