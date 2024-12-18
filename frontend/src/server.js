import React from 'react';
//import { useMediaQuery } from '@mui/material';
import {
    List,
    Datagrid,
    TextField,
    EditButton,
    Create,
    SimpleForm,
    TextInput,
    Edit,
    DeleteButton,
    ShowButton,
    Show,
    SimpleShowLayout,
    ReferenceManyField,
    CreateButton,
} from 'react-admin';
//import {Typography} from "@mui/material/colors";
export const server = (props) => (
<List {...props}>
    <Datagrid rowClick="show">
        <TextField source="id" />
       <TextField source="ip" />
       <TextField source="user_ssh" />
        <TextField source="password_ssh" />
        <TextField source="distro_installata" />
        <TextField source="lista_programmi" />
        <ShowButton/>
        <DeleteButton/>
        <EditButton/>
   </Datagrid>
</List>
);
export const serverCreate=()=>(
    <Create>
        <SimpleForm>
            <TextInput source="id"/>
            <TextInput source="user_ssh" />
            <TextInput multiline source="password_ssh" />
        </SimpleForm>
    </Create>
);
export const serverEdit=()=>(
    <Edit>
        <SimpleForm>
            <TextInput  source="user_ssh"/>
            <TextInput  source="password_ssh"/>
        </SimpleForm>
    </Edit>
)
export const swarmShow=()=>(
    <Show >
        <SimpleShowLayout>
            <TextField source="id" />
            <TextField source="ip"/>
            <TextField source="user_ssh"/>
            <TextField source="password_ssh"/>
            <TextField source="distro_installata"/>
            <TextField source="lista_programmi"/>
            <ReferenceManyField label="nuovo" reference="swarm" target="server_id">
                <CreateButton/>
                <Datagrid>
                    <TextField source="uuid" />
                    <TextField source="nome" />
                    <TextField source="file_content_text" />
                    <TextField source="server_id" />
                </Datagrid>
            </ReferenceManyField>
        </SimpleShowLayout>
    </Show>
)
