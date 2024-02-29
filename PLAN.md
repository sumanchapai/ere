Event reminder.

## `ere daily`

`ere daily` checks the daily updates and prints any daily update for today if
today's updates haven't been checked before. The algorithm goes something like
this:

- Get the english date for today in a given format.
- Checks the ~/.ere/checked.txt to see if the last line contains today's date.
- If it does:
  - Do nothing as today's update has already been checked.
- It it doesn't:
  - Run the `ere check` command with today's date in a given format. For
    example: `ere check 2022-10-29`.
  - Append today's date to ~/.ere/checked.txt file

This command is supposed to be put in user's ~/.zshrc file so that terminal
dwellers will get the update today's reminders the first time they open
terminal.

## `ere check <date>`

`ere check <date>` takes an optional date argument and prints the updates for
that date. If no date argument is provided, checks the update for today. See the
information below for format for <date>.

### Date format

Note that the legal date formats are "YYYY-MM-DD$xx", "MM-DD$xx" where "xx" is
supposed to be calendar prefixes (BS for bikram sambat and AD for Julian
Calendar). Note that "MM$xx" and "DD$xx" are both valid formats as well, but
which one is meant day or month depends on the context at which the date command
is used. When used with the `check` command only `MM$xx` is inferred and when
used with the `add` command only `DD$xx` is inferred.

## `ere add <date>`

Adds a reminder on a given date. Check [date format](###date-format).

#### One off events

To be reminded about poster presentation coming up on 29th october, run\
`ere add 2025-10-29$AD --name "poster presentaion"`.

To also be reminded the poster presentation one day before the presentation date
run\ `ere add 2025-10-29$AD --name "poster presentaion" --knock="1"`.

To be reminded the poster presentation every day for one week before the
presentation run\
`ere add 2025-10-29$AD --name "poster presentaion" --knock="1,2,3,4,5,6,7"`.

#### Yearly events (like birthdays)

To be reminded about your friend's birthday on Kartik 12, run\
`ere add 7-12$BS --name "suman's birthday" --knock="1"`.

Note that we don't mention year in the date format because that would make it a
one off event. We want to be reminded every year on our friend's birthday.

#### Monthly events (like grocery at the beginning of the month)

To be reminded about Sakranti (first day of the month in the Nepali calendar),
run\
`ere add 1$BS --name "sakranti"`.

To be reminded about taking the trash out on the 7th of the month. run\
`ere add 7$AD --name "take the trash out"`.

To be reminded one day before taking the trash out. run\
`ere add 7$AD --name "take the trash out" --knock="1"`.

#### Events files:

These are the files where the information about the event will be kept.

- weekly.json
- monthly.json
- oneoff.json

### How does the `ere check` work ?

1. First gets today's date in AD.
1. Gets today's day in BS, by converting the AD date to BS.
1. Goes through the event files to see if there's an event for today. This means
   both today's event and a future event that is to be knocked about.

### `ere delete <id>`

Delete the event with given <id>. Note that you can view the id of events if you
run `ere search`

### `ere search <search-string>`

List the events whose name matches the search-string. The search string is a
regular expression.
