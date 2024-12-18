import * as React from "react";
import { List, Datagrid, TextField,SimpleForm,Create,TextInput,EditButton,Edit,DeleteButton,ReferenceInput,SelectInput} from 'react-admin';
export const swarm = () => (
    <List>
        <Datagrid>
            <TextField source="uuid" />
            <TextField source="nome" />
            <TextField source="file_content_text" />
            <TextField source="server_id" />
            <DeleteButton/>
            <EditButton/>
        </Datagrid>
    </List>
);
export const swarmCreate=()=>(
    <Create>
        <SimpleForm>
            <TextInput source="uuid" />
            <TextInput multiline source="nome" />
            <TextInput source="file_content_text" />
            <ReferenceInput source="server_id" reference="server" target="id" >
                <SelectInput  optionText="user_ssh" optionValue="id" target="id" source="server"/>
            </ReferenceInput>
        </SimpleForm>
    </Create>
);
export const swarmEdit =(props)=> (
    <Edit {...props}>
        <SimpleForm>
            <TextInput source="uuid" />
            <TextInput source="nome" />
            <TextInput source="file_content_text" />
        </SimpleForm>
    </Edit>
);//sta compilando ma lentamentissimo