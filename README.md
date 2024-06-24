# SRE Technical challenge - HyperAutomation

- Task:
    - Saturday is a weekly once planned service day
    - SRE -> CMT (Central monitoring team) every week
    - Automate this task
        - create script
        - use JJ to run the script weekly once
    - preferred language (Go or Python)
    - I/P:
        - Email recipients
        - Email sender
        - Email message
        - Email server configurations
- Goal: Create a Jenkins pipeline that would run this script every week on Saturday morning. The pipeline should take the above-mentioned script inputs as parameters.

- P.S:
    - Above mentioned script resides in some GitHub repository (not public). So, the Jenkins job can clone that repository and run the script