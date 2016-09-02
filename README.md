# credit-card
Basic credit card account processing

Input should be a file where each line is one out of three types of commands:
-- `Add [account holder's name] [credit card number] [credit limit]`
   -- Initializes an account with a balance of zero.
-- `Charge [account holder's name] [amount]`
   -- Charges (adds) amount to the specified account
   -- Does nothing if adding the amount would cause the balance to exceed the account's limit
   -- Does nothing if account is invalid
-- `Credit [account holder's name] [amount]`
   -- Credits (subtracts) amount to the specified account
   -- Does nothing if account is invalid

Output is an alphabetical list of accounts, of the format `[Name]: [Balance]`, or `[Name]: error` if the credit card number was invalid.

Credit cards are considered valid if they are fewer than 19 characters, all numeric, and pass Luhn 10 validation.

All input is assumed to be well-formed, for the purposes of this project.
