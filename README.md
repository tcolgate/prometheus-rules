# A collection of useful prometheus rules

This is intended as a home for collections of useful, reusable prometheus rules.

 - *time.rules*: Rules for building time/date awareness. These are approximations, with some room for improvement:
  - time\_current\_year: The current year
  - time\_start\_current\_year\_seconds: Start of current year, second since Unix Epock
  - time\_end\_current\_year\_seconds:  End of current year, second since Unix Epock
  - time\_end\_of\_current\_day\_seconds:  End of current day, second since Unix Epock
  - time\_current\_day\_remaining\_seconds: Seconds left in the current day
  - time\_is\_leap\_year\_bool: 1 indicates that you are currently in a leap year
  - time\_end\_of\_current\_[month]\_seconds: month is a three letter month indicator
  - time\_current\_month: 1 == Jan, 12 == December
  - time\_end\_of\_current\_month\_seconds: End of the current month, seconds since Unix Epoch
  - time\_current\_month\_remaining\_seconds: Seconds remaining in the current month
  - time\_current\_weekday: The current week day 1 is Monday, 7 is Sunday
  - time\_is\_weekday\_bool: Check if the current day is a weekday (not Saturday or Sunday)

# License

These configs are released under the Apache 2.0 license. All images
downloaded are subject to their individual licenses.
