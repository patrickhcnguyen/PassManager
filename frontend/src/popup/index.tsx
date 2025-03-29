import { createRoot } from 'react-dom/client'
import Popup from './popup'
import '../index.css'

createRoot(document.getElementById('popup-root')!).render(
  <Popup />
) 