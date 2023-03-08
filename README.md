# Informabot

A Telegram bot to assist other undergraduate Computer Science students at the
University of Bologna.

## Dependencies

[Node.js](https://www.nodejs.dev) 12.2.2 or higher is required.

## Running

To run the bot from the project directory, simply use

```bash
npm start -- $TOKEN
```

where `$TOKEN` is the authorization token you got from
[BotFather](https://core.telegram.org/bots#6-botfather).

## Commands

When setting up your bot with
[BotFather](https://core.telegram.org/bots#6-botfather), you'll be
asked to "send a list of commands for your bot". Consider using this one:

> help - Elenca i comandi disponibili
>
> appunti - Appunti su Notion
>
> aggiorna - Aggiorna le azioni del bot
>
> cercogruppo - Cerca un gruppo di progetto
>
> dns - I DNS di ALMAWIFI
>
> admstaff - Tutte le info di AdmStaff
>
> gruppi - Gruppi Telegram, server Discord e subreddit
>
> lezionioggi - Orari lezioni di oggi (tuo anno)
>
> lezionioggi1 - Orari lezioni di oggi (1° anno)
>
> lezionioggi2 - Orari lezioni di oggi (2° anno)
>
> lezionioggi3 - Orari lezioni di oggi (3° anno)
>
> lezionidomani - Orari lezioni di domani (tuo anno)
>
> lezionidomani1 - Orari lezioni di domani (1° anno)
>
> lezionidomani2 - Orari lezioni di domani (2° anno)
>
> lezionidomani3 - Orari lezioni di domani (3° anno)
>
> materiali - Libri, materiali, prove e altro su CSUnibo
>
> noncercogruppo - Smetti di cercare un gruppo di progetto
>
> registrate - Lezioni registrate su OneDrive

> ludopatico - Tenta la fortuna
>
> scelta - Elenchi esami a scelta
>
> stickers - Pacchetti degli adesivi Telegram del corso
>
> tesi - Tesi DISI proposte e assegnate
>
> coursedesc - Documento con tutte le descrizioni dei corsi in inglese
>
> tirocinio - Strumento per riassumere le attività svolte durante il proprio tirocinio
>
> architettura - Tutto su Architettura degli elaboratori
>
> logica - Tutto su Logica per l'informatica
>
> programmazione - Tutto su Programmazione
>
> algebra - Tutto su Algebra e geometria
>
> algoritmi - Tutto su Algoritmi e strutture di dati
>
> analisi - Tutto su Analisi matematica
>
> calcolo - Tutto su Calcolo numerico
>
> ottimizzazione - Tutto su Ottimizzazione combinatoria
>
> linguaggi - Tutto su Linguaggi di programmazione
>
> reti - Tutto su Reti di calcolatori
>
> tecnologie - Tutto su Tecnologie web
>
> probabilita - Tutto su calcolo delle probabilità e statistica
>
> sistemi - Tutto su Sistemi operativi
>
> basi - Tutto su Basi di dati
>
> ingegneria - Tutto su Ingegneria del software
>
> apprendimento - Tutto su Introduzione all'apprendimento automatico
>
> cybersecurity - Tutto su Fondamenti di cybersecurity
>
> teorica - Tutto su Informatica teorica
>
> progetto - Tutto su Progetto di sistemi virtuali
>
> fisica - Tutto su Fisica
>
> applicazioni - Tutto su Laboratorio di applicazioni mobili
>
> storia - Tutto su Storia dell'informatica e dei dispositivi di calcolo
>
> strategia - Tutto su Strategia aziendale

## Adding new commands

### Actions

Actions are stored in `json/actions.json`. Each key is the name of the command
triggering the action, while each value is an object describing the action
itself through several attributes:

- `type` specifies the command logic. Available types are listed below, each
  with its own specialized attributes;
- `description` (optional) sums up the objective of the action.

#### `message`

The bot replies with a static message, specified by the `text` attribute.

#### `list`

The bot replies with an automatically generated list, preceeded by the `header`
attribute. Each element in the list is generated substituting the placeholders
in the `template` attribute with the elements of a different array from the
`items` attribute matrix.

#### `help`

The bot replies listing each command-description pair. If a command has no
description, it is not listed.

#### `luck`

Tests your luck.

#### `alias`

These commands are just aliases for others. The `command` attribute specifies
which command this alias is referring to. Beware of circular alias chains, which
will result in a stack overflow.

#### `lookingFor`

The bot adds the user to the list of people looking for project mates in this
chat, and replies with the updated list. `singularText`, `pluralText`, and
`chatError` attributes are used as custom messages to communicate with the user.

#### `notLookingFor`

The bot removes the user from the list of people looking for project mates in
this chat. `chatError` and `notFoundError` messages are used to communicate with
the user.

#### `yearly`

Much like an alias with `abc` as its `command` attribute value runs the `abc`
action, a `yearly` action with `abc` as its `command` attribute value may run
either the `abc1`, `abc2`, or the `abc3` command, depending on the chat. The
bot inspects the chat title, and attempts to figure out the appropriate year.
If the bot can't figure out the year, the `noYear` attribute value is used as
a default reply.

#### `todayLectures`

Scrapes today's timetable from `url`, using `title` as header. On faliure,
`fallbackText` is used as a reply.

#### `tomorrowLectures`

Scrapes tomorrow's timetable from `url`, using `title` as header. On faliure,
`fallbackText` is used as a reply.

#### `course`

Puts together a summary for a given course, featuring its `name`, Virtuale link
(`virtuale`), Teams link (`teams`), official `website` link, as well as the
email addresses of the `professors`.

#### `update`

Pulls from the original repo and reloads all of the JSON files, without
restarting the bot. It can only be run by administrators in a general group.
If it is not run in a general group, a `noYear` message is sent. Otherwise, if
it is not run by an administrator, a `noMod` message is sent. Otherwise the
update is attempted, and the `started` message is sent. At the end of the
operation, depending on the outcome, either the `ended` or the `failed` message
may be sent.

### Memes

Memes are stored in `json/memes.json`. Each key is the name of the command
triggering the meme, while each value is the content of the reply by the bot (a
simple static message).

## Acknowledgments

Many thanks to [@Wifino](https://github.com/Wifino), who wrote the original
codebase.
