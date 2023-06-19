import { NitroApp } from 'nitropack';
import { Auth } from './defineAuth';

export interface ExtendedNitroApp extends NitroApp {
  auth: Auth;
}

export default function (app: NitroApp = useNitroApp()) {
  return app as ExtendedNitroApp;
}
