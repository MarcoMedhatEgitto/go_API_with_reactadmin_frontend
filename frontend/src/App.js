import React from 'react';
import { Admin, Resource,EditGuesser,defaultTheme} from 'react-admin'
import {server, serverCreate, swarmShow} from './server';
import {swarm, swarmCreate,swarmEdit} from './swarm';
import { createTheme } from '@mui/material/styles';
import dataProvider from './dataprovider';
import serv from '@mui/icons-material/Computer';
import swar from '@mui/icons-material/GroupWork';
const lightTheme = defaultTheme;
const darkTheme = createTheme({
    palette: {
        mode: 'dark',
        primary: {
            main: '#2196f3',
        },
        secondary: {
            main: '#2196f3',
        },
        text: {
            primary: '#2196f3',
            secondary: '#2196f3',
        },
    },
})
const App = () => (
        <Admin   theme={lightTheme}
                 darkTheme={darkTheme}
                 dataProvider={dataProvider}
        >
    <Resource name="server" list={server} edit={EditGuesser} create={serverCreate} show={swarmShow} icon={serv}/>
    <Resource name="swarm" list={swarm} create={swarmCreate} edit={swarmEdit} icon={swar}/>
    </Admin>
);
export default App;