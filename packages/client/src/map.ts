import '@utils/map/here_wego';
import { HereWeGO } from '@utils/map/here_wego';
import { $ } from '@utils/dom';

setTimeout(() => {
  let $elem = $('here-wego');
  $elem.set("test", "here-wego after 2 seconds");
}, 2000);

setTimeout(() => {
  let $elem = $('here-wego');
  let elem = $elem[0] as HereWeGO;
  elem.test = "here-wego after 4 seconds"
}, 4000);

