import '@utils/map/here_wego';


setTimeout(() => {
  let $elem = document.querySelector('here-wego')
  console.log($elem);
  $elem?.setAttribute("test", "here-wego after 2 seconds")
}, 2000);

