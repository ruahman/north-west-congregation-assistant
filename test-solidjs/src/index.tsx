/* @refresh reload */
import { render } from 'solid-js/web'
import  { defineCustomElement } from 'test-stencil/dist/components/my-component'
defineCustomElement();

import './index.css'
import App from './App'

const root = document.getElementById('root')

render(() => <App />, root!)
