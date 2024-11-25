import { HereMap } from '@utils/map/here';
import { $ } from '@utils/dom';

// this creates here-map tag
new HereMap();

setTimeout(() => {
  let $elem = $('here-map');
  $elem.set("test", "here-wego after 2 seconds");
}, 2000);

setTimeout(() => {
  let $elem = $('here-map');
  let map = $elem[0] as HereMap;
  map.test = "here-wego after 4 seconds"
}, 4000);

