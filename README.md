This is an event reminder CLI. It currently supports two calendar systems
Gregorian (A.D) and Nepali (B.S).

## Adding events

#### One off events

To be reminded about calling mom tomorrow, run\
`ere add +1 --title "call mom"`.

To be reminded about poster presentation coming up on 29th october, run\
`ere add 2025-10-29-AD --title "poster presentation"`.

To also be reminded the poster presentation one day before the presentation date
run\
`ere add 2025-10-29-AD --title "poster presentation" --knock="1"`.

To be reminded the poster presentation every day for one week before the
presentation run\
`ere add 2025-10-29-AD --title "poster presentation" --knock="1,2,3,4,5,6,7"`.

#### Yearly events (like birthdays)

To be reminded about your friend's birthday on Kartik 12, run\
`ere add "*-7-12-BS" --title "suman's birthday" --knock="1"`.

Note that we don't mention year in the date format because that would make it a
one off event. We want to be reminded every year on our friend's birthday.

#### Monthly events (like grocery at the beginning of the month)

To be reminded about Sakranti (first day of the month in the Nepali calendar),
run\
`ere add "*-*-1-BS" --title "sakranti"`.

To be reminded about taking the trash out on the 7th of the month. run\
`ere add "*-*-7-AD" --title "take the trash out"`.

To be reminded also one day before taking the trash out. run\
`ere add "*-*-7-AD" --title "take the trash out" --knock="1"`.

## Viewing events

Run `ere ls` to view events. You can list events whose title matches a given
regex by runing `ere search <pattern>`.

For example, to list events that have the word birthday in their title, run\
`ere search 'birthday'`

## Delete event

To delete an event by given id, run\
`ere delete <id>`

## Checking for events

To view today's events, run\
`ere check`

To view yesterday's events, (note the extra -- because a - alone means flag is
specified) run\
`ere check -- -1`

To view day after tomorrow's events, run\
`ere check 2`

To view events on Oct 29, 2024, run\
`ere check 2024-10-29-AD`

To view events on Baiskah 10, 2078, run\
`ere check 2078-1-10-BS`

## Lookahead

To look `5` (use a number of your choice) days into the future, run\
`ere la 5`

Note that this computation takes time on the order of the number of days.

### Automatically checking for today's events

There's is also a command called `ere daily` that runs the `ere check`
internally if `ere check` hasn't been run even once today. I have placed the
`ere daily` at the end of my `~/.zshrc` file so that when I open the terminal
for the first time on any given day, I see the events for that day. The next
time I open terminal on that day, it doesn't show anything. But of course, I
could just run `ere check` anytime to see events for that day or any day.

## Syncing events between two computers.

All the data files are stored locally in `.ere` folder in the home directory.
Run `ls ~/.ere` to see these files. The list of events is stored in a file
called `~/.ere/events.toml`. I have personally created this file in my dotfiles
directory and created a symlink to `~/.ere/events.toml` by running:

```
mkdir -p ~/dotfiles/ere
ln -sf "${HOME}/dotfiles/ere/events.toml" "${HOME}/.ere/events.toml"
```

Note that my dotfiles folder lives in the home directory.

Because my dotfiles directory is git version controlled, when I install `ere` on
another computer and pull my dotfiles directory (and create the symlink like
mentioned above), I have the same list of events available on both computers.
Note though that, you probably don't want to do this if your dotfiles directory
is public because this would make your events public.

## Archive old events

To archive past events, run `ere archive`. This goes through each event and
checks if this event won't arrive in the future again. If such, then it puts
this event in the `archived.toml` file. You can list all archived events using
`ere ls --archive` and search through them using
`ere search <pattern> --archive`

## Installing / Updating `ere`

### For macOS/Linux

```sh
rm -f `which ere`
curl -s https://ere-sumanchapai.vercel.app/install.txt | sh
```

### For Windows

Have go >= 1.21 installed. Clone the repo, then run `go install .`
