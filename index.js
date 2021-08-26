/*
 * Stefano Volpe
 * 08/26/2021
 *
 * index.js: entry point of the project
 */

import axios from 'axios';
import TelegramBot from 'node-telegram-bot-api';

// Setup
if (process.argv.length != 3)
  console.log("usage: node index.js token");
process.env.NTBA_FIX_319 = 1;
const bot = new TelegramBot(process.argv[2], {polling: true});

const options = {
  "parse_mode": "HTML",
  "disable_web_page_preview": 1,
};

function timetable() {
  axios
    .get("https://corsi.unibo.it/laurea/informatica/orario-lezioni/@@orario_reale_json?anno=2")
    .then(res => {
      let now = new Date();
      let todayLectures = [];
      for (let i = 0; i < res.data.length; ++i) {
        let start = new Date(res.data[i].start);
        if (start.getFullYear() === now.getFullYear() &&
          start.getMonth() === now.getMonth() &&
          (start.getDate() === now.getDate()))
          todayLectures.push(res.data[i]);
      }

      let text = "";
      todayLectures.sort((a, b) => {
        if (a.start > b.start)
          return 1;
        if (a.start < b.start)
          return -1;
        return 0;
      });
      for (let i = 0; i < todayLectures.length; ++i)
        text += "*" + todayLectures[i].title + "* " + todayLectures[i].time + "\n";
      if (todayLectures.length !== 0)
        bot.sendMessage(msg.chat.id, text, {parse_mode: "markdown"}).catch(e => console.error(e.stack));
      else
        bot.sendMessage(msg.chat.id, "Non ci sono lezioni oggi. SMETTILA DI PRESSARMI", {parse_mode: "markdown"}).catch(e => console.error(e.stack));
    }).catch(e => console.error(e.stack));
}

function course(name, virtuale, teams, website, emails) {
  // Remember to double up every % you want to escape!
  bot.sendMessage(msg.chat.id, String.format("<b>%s</b>\n<a href='https://vir" +
    "tuale.unibo.it/course/view.php?id=%s'>Virtuale<a>\n<a href='https://team" +
    "s.microsoft.com/l/meetup-join/19%%3ameeting_%s%%40thread.v2/0?context=%%" +
    "7b%%22Tid%%22%%3a%%22e99647dc-1b08-454a-bf8c-699181b389ab%%22%%2c%%22Oid" +
    "%%22%%3a%%22080683d2-51aa-4842-aa73-291a43203f71%%22%%7d'>Videolezione</" +
    "a>\n<a href='https://www.unibo.it/it/didattica/insegnamenti/insegnamento" +
    "/%s'>Sito</a>\n<a href='https://www.unibo.it/it/didattica/insegnamenti/i" +
    "nsegnamento/%s/orariolezioni'>Orario</a>\n%s", name, virtuale, teams,
    website, website, emails.join("@unibo.it<br>"), options));
}

function answer(msg, str) {
  bot.sendMessage(msg.chat.id, str, options);
}

function onMessage(msg) {
  if (msg.text)
    switch (msg.text.toString().toLowerCase().split(" ")[0]) {
      // Generals
      case "/appunti":
        bot.sendMessage(msg.chat.id, "Grazie ai nostri gentili contribuenti, ecco l'elenco dei Notion:\n\n➽<a href='https://www.notion.so/Algebra-e-geometria-00d4b98a5d974879aaf39457ede3261a'><b>Algebra e Geometria (Luizo)</b></a>\n➽<a href='https://www.notion.so/Algebra-a65a99336ccc499ead0637365a3bd0cd'><b>Algebra e Geometria (Alex)</b></a>\n\n➽<a href='https://www.notion.so/Algoritmi-e-strutture-dati-70a01e43fa034859bb0c8cd6d744e6d6'><b>Algoritmi (Alex)</b></a>\n➽<a href='https://www.notion.so/Algoritmi-e-Strutture-di-Dati-da9a9d634c6f433cb778cdd02bead894'><b>Algoritmi (Luizo)</b></a>\n\n➽<a href='https://www.notion.so/Analisi-1895389f8b9a465e98f2a868fc917c53'><b>Analisi (Alex)</b></a>\n➽<a href='https://www.notion.so/Analisi-Prova-unica-ab60229e9ac5455cb69b24b3c41fd0b1'><b>Analisi (Esercizi) (Fabrizio)</b></a>\n\n\n➽<a href='https://www.notion.so/Logica-logico-1adfde3168d94cc5ac461da479d113ee'><b>Logica (Alex)</b></a>\n➽<a href='https://www.notion.so/Preparazione-logica-3-CFU-8bf160d661d149f9939d5a48e72edf05'><b>Logica (Modulo da 3 CFU)(Alex)</b></a>\n➽<a href='https://www.notion.so/Ripasso-bc03206bfa034bed8f3f521778a61254'><b>Logica (Andrea)</b></a>\n\n➽<a href='https://www.notion.so/Appunti-784f6703da1447028ea95a52eda74f38'><b>Programmazione (Andrea)</b></a>\n\n<b>N.B. Ai sensi delle Leggi Infernali del Sommo CSC, i relatori non hanno alcuna responsabilità riguardo eventuali informazioni errate presenti all'interno degli appunti.</b>", option);
        break;
      case "/lezionidioggi":
        timetable();
        break;
      case "/libri":
        bot.sendMessage(msg.chat.id, "<b>LIBRI </b>\n<a href='https://liveunibo-my.sharepoint.com/:f:/g/personal/gurjyot_wanga_studio_unibo_it/Ej2rZIDZqklPpin8votjqR0ByrhNTUQW7qpklJjeLS196w?e=cbBVME'><b> LIBRI DI TUTTE LE MATERIE</b></a>", option);
        break;
      case "/link":
        bot.sendMessage(msg.chat.id, "<b>GRUPPI DELLE MATERIE</b>\n<a href='https://t.me/joinchat/2hUcGLfY7Gc2MTA8'><b>➽ Sistemi operativi </b></a>\n<a href='https://t.me/joinchat/4v-08oT6QWk0ZmM8'><b>➽ Linguaggi di programmazione </b></a>\n<a href='https://t.me/joinchat/Sw2Ykp0-0mM1Zjdk'><b>➽ Reti di calcolatori </b></a>\n<a href='https://t.me/joinchat/h1lypfBFdEZkYzFk'><b>➽ Calcolo Numerico </b></a>\n<a href='https://t.me/joinchat/klw79l5tkPM1NWQ0'><b>➽ Ottimizzazione Combinatoria </b></a>\n\n<a href='https://discord.gg/YcAc2rdS3H'><b>☻ DISCORD ☻</b></a>\n\n<b>Libri, Materiali e Registrazioni</b>\n<a href='https://liveunibo-my.sharepoint.com/:f:/g/personal/gurjyot_wanga_studio_unibo_it/EnTEAPe1X-RHoisCwNfQykQBWGOXHfwEqSdQcOqCWsQFgw?e=SYwCR7'><b>Primo Anno</b></a>\n\n<a href='https://liveunibo-my.sharepoint.com/:f:/g/personal/gurjyot_wanga_studio_unibo_it/EkH1O5CfQk9FniJopixNv0YBWWtW-GooDFuSx_9kbgOF1Q?e=RX0Gzx'><b>Secondo Anno</b></a>\n\n<a href='https://liveunibo-my.sharepoint.com/:f:/g/personal/simone_folli2_studio_unibo_it/Ep7wMjaQIeJGlM7vRd5T96cBf-odnowMZYahxYdPKyP1-g'><b>Caricare File Su Onedrive </b></a>", option);
        break;
      case "/orari":
        bot.sendMessage(msg.chat.id, "<b>ORARI 2° SEMESTRE </b>\n\n<a href='https://www.unibo.it/it/didattica/insegnamenti/insegnamento/2020/366975/orariolezioni'><b>ALGEBRA E GEOMETRIA</b></a>\n<a href='https://www.unibo.it/it/didattica/insegnamenti/insegnamento/2020/350957/orariolezioni#412691'><b>ALGORITMI E STRUTTURE DATI (MOD 1)</b></a>\n<a href='https://www.unibo.it/it/didattica/insegnamenti/insegnamento/2020/320573/orariolezioni#447976'><b>ANALISI (MOD 2)</b></a> ", option);
        break;
      case "/registrate":
        bot.sendMessage(msg.chat.id, "<b>Qui le videolezioni registrate</b>\n\n<a href='https://liveunibo-my.sharepoint.com/:f:/g/personal/simone_folli2_studio_unibo_it/Ep7wMjaQIeJGlM7vRd5T96cBhoND3NZL_y1hQ7BpNHy-_Q'><b>LEZIONI REGISTRATE</b></a>\n<a href='https://liveunibo-my.sharepoint.com/:f:/g/personal/simone_folli2_studio_unibo_it/Ep7wMjaQIeJGlM7vRd5T96cBf-odnowMZYahxYdPKyP1-g'><b>CARICARE LEZIONI SU ONEDRIVE</b></a>\n", option);
        break;
      // First year
      case "/architettura":
        course("Architettura degli elaboratori", "18282", "ZjM2MGUxNTAtODA1NC00N2NiLWEwOWMtYTllMzZkOGQ0MjMx", "2020/350960", ["ivan.lanese"]);
        break;
      case "/logica":
        course("Logica per l'informatica", "21407", "YTNjMjI0NzctNzU2OC00MWI3LTlkNDctMTcwZDg4OGVjNjRk", "2020/455095", ["claudio.sacerdoticoen"]);
        break;
      case "/programmazione":
        course("Programmazione", "17653", "OTQyM2U2MzEtNjc3NS00N2ZmLWJlOTgtOGMzM2JmMGJhNDA4", "2020/320574", ["cosimo.laneve"]);
        break;
      case "/algebra":
        course("Algebra e geometria", "17870", "MGU2ZWEyNjgtYThmZi00ZTMyLTg4YWUtZTAwZDViZTY1Nzkw", "2020/366975", ["marta.morigi"]);
        break;
      case "/algoritmi":
        course("Algoritmi e strutture di dati", "20930", "NDJjMTA4ZGEtODMzNy00NjZmLThhNmYtMmUzYWU4YzhiMjVl", "2020/350957", ["gianluigi.zavattaro", "pietro.dilena"]);
        break;
      case "/analisi":
        course("Analisi matematica", "18045", "ODk1NjI3MGMtZThhOC00MmU4LTljYmQtOWNlZDdiYjhhYjhk", "2020/320573", ["marco.mughetti", "daniele.morbidelli"]);
        break;
      // TODO: Second year
      case "/calcolo":
        course("Calcolo numerico", "", "", "2021/320581", ["elena.loli"]);
        break;
      case "/ottimizzazione":
        course("Ottimizzazione combinatoria", "", "", "2021/460495", ["ugo.dallago"]);
        break;
      case "/linguaggi":
        course("Linguaggi di programmazione", "", "", "2021/320579", ["roberto.gorrieri", "maurizio.gabbrielli", "saverio.giallorenzo2"]);
        break;
      case "/reti":
        course("Reti di calcolatori", "", "", "2021/455456", ["luciano.bononi"]);
        break;
      case "/sistemi":
        course("Sistemi operativi", "", "", "2021/320578", ["renzo.davoli"]);
        break;
      // Memes
      case "/alice":
        answer(msg, "<b>@alii_benatti, registri la lezione di oggi?</b>");
        break;
      case "/altribot":
        answer(msg, "<b>VAFFANCULO ALICE ED ANCHE AGLI ALTRI BOT</b>");
        break;
      case "/bestmod":
        answer(msg, "<b>SICURAMENTE NON LUIZO.</b>");
        break;
      case "/betto":
        answer(msg, "<b>S I M P</b>");
        break;
      case "/biagio":
        answer(msg, "<b>Biagio TVB</b>");
        break;
      case "/chiara":
        answer(msg, "<b>yo te rao!</b>");
        break;
      case "/csc":
        answer(msg, "<b>In nomine Dei Nostri Luciferi Excelsi Ghepardi CSC;\nNel Nome di Claudio Sacerdoti Coen\nDominatore della logica,\nVero meta-Dio,\nOnnipotente e Ineffabile,\nColui che creò l’ uomo\na sua meta-immagine e meta-somiglianza.\nIo invoco\nle Forze di CSC\naffinché infondano\nil loro potere infernale in me.</b>");
        break;
      case "/domande":
        answer(msg, "<b>@gabboTRNGL MANCA POCO ALLA FINE DELLA LEZIONE. VEDI DI STARE ZITTO! TU NON HAI DOMANDE!</b>");
        break;
      case "/flamealice":
        answer(msg, "<b>ALICE FAI SCHIFO!</b>");
        break;
      case "/foxy":
        answer(msg, "<b>FOXY SEI BELLISSIMO</b>");
        break;
      case "/giuseppe":
        answer(msg, "<b>Vuoi vedere i miei cyberPiedini?</b>");
        break;
      case "/hokage":
        answer(msg, "<b>Matteo Manuelli, [09.03.21 22:28]</b>\nsemplicemente vi straccerò a mnk game.\nGià pregusto la faccia soddisfatta di zavattarro\n\n<b>Niccolò CEO dei dissing e del flame, [09.03.21 22:29]</b>\nVai bro, sarai Hokage");
        break;
      case "/laneve":
        answer(msg, "<b>nCi sono dei bug!\nF A N T A S T I C O!</b>");
        break;
      case "/luiso":
        answer(msg, "<b>LUIZO TVB :></b>");
        break;
      case "/luizo":
        answer(msg, "<b>LUIZO HAI ROTTO LE PALLE!</b>");
        break;
      case "/nobel":
        answer(msg, "<b>Ho sempre creduto in Matteo Manuelli. Lui ha scritto il sacro algoritmo.</b>");
        break;
      case "/ping":
        answer(msg, "<b>PONG</b>");
        break;
      case "/rinunciaaglistudi":
        answer(msg, "<b>Lascia stare, non fa per te.</b>");
        break;
      case "/zavattarro":
        answer(msg, "<b>Zavattarro sarà fiero di me (Cit. M.M.)</b>");
    }
}

bot.on('message', onMessage);
