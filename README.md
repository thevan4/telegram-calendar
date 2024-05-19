# Motivation

Everything I've seen about date selection in Telegram is done quite badly.

So I wanted to make a beautiful, convenient and universal implementation.

# Demo

## Gif

![](https://github.com/thevan4/telegram-calendar-examples/blob/main/demo/tgc-demo-1.gif)

## Live demo

[Bot](https://t.me/buttons_calendar_example_bot)

# Settings

## Values

Default values are given at the end of the line.

- YearsBackForChoose(int) - how many years in the past are available for selection when opening a calendar with year selection. ["0"]
- YearsForwardForChoose(int) - how many years in the future are available for selection when opening a calendar with year selection. ["3"]
- DaysNames([7]string) - names of the days of the week (the week always starts on Monday). ["Mo", "Tu", "We", "Th", "Fr", "Sa", "Su"]]
- MonthNames([12]string) - month names. ["Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"]]
- HomeButtonForBeauty(string) - the icon of the button which, when clicked, goes to the current month of the user. ["🏩"]
- PrefixForCurrentDay(string) - prefix for the current day. [""]
- PostfixForCurrentDay(string) - postfix for the current day. ["🗓"]
- PrefixForNonSelectedDay(string) - prefix for the day that is not available for selection. [""]
- PostfixForNonSelectedDay(string) - postfix for a day that is not available for selection. ["❌"]
- PrefixForPickDay(string) - prefix for a day that is available for selection. [""]
- PostfixForPickDay(string) - postfix for the day that is available for selection. [""]
- UnselectableDaysBeforeTime(time.Time) - all dates specified before this time (exactly time, not date!) will be unavailable. ["01.01.2023 UTC"].
- UnselectableDaysAfterTime(time.Time) - all dates specified after this time (exactly time, not date!) will be unavailable. ["01.01.2030 UTC"]]
- UnselectableDays(map[time.Time]struct{}) - map unavailable dates/days. [""]
- Timezone(time.Location) - your timezone. ["UTC"]

## About timezones

All incoming requests with time are converted to the originally specified timezone. That is, the timezone of the input (user) will be converted to the specified timezone.
For example: utc+0 timezone is set in the generator, the user came with current date utc+3. The calculation of unavailable days will be based on the generator data, but the calculation of the current day will be based on the user's time.
That is, if you set "all days starting from the last day are unavailable" (in Go it is *.AddDate(0, 0, -1)), the last day will be marked as unavailable for the generator. But for the user this "last day" can be considered as the current day (if the local time is 21:00-23:59).
This behavior is considered normal and is not a bug.

# Examples

Examples [here](https://github.com/thevan4/telegram-calendar-examples)

All in all the project is ready to integrate [not only](https://github.com/thevan4/telegram-calendar-examples/tree/main/standalone_service) with golang.