import { fetchUtils } from 'react-admin';
const apiUrl = 'http://localhost:8080';
const httpClient = fetchUtils.fetchJson;
const dataProvider = {
    getList: (resource, params) => {
        const url = `${apiUrl}/${resource}`;
        return httpClient(url).then(({ json }) => ({
            data: json,
            total: json.length,
        }));
    },
    create: (resource, params) => {
        // make a POST request to the backend API
        const url = `${apiUrl}/${resource}`;
        const options = {
            method: 'POST',
            body: JSON.stringify(params.data),
            headers: new Headers({ 'Content-Type': 'application/json' }),
        };

        return httpClient(url, options).then(({json})=>({
            data:json,
        }));
    },
    getOne: (resource, params) => {
        // make a POST request to the backend API
        const url = `${apiUrl}/${resource}/${params.id}`;
        const options = {
            method: 'GET',
            //body: JSON.stringify(params.id),
            headers: new Headers({'Content-Type': 'application/json'}),
        };

        return httpClient(url, options).then(({json})=>({
            data:json,
        }));
    },
    update: (resource, params) => {
        // make a POST request to the backend API
        const url = `${apiUrl}/${resource}/${params.id}`;
        const options = {
            method: 'PUT',
            body:JSON.stringify(params.data),
            headers: new Headers({ 'Content-Type': 'application/json' }),
        };

        return httpClient(url, options).then(({json})=>({
            data:json,
        }));
    },
    delete: (resource, params) => {
        // make a POST request to the backend API
        const url = `${apiUrl}/${resource}/${params.id}`;
        const options = {
            method: 'DELETE',
            body: JSON.stringify(params.data),
            headers: new Headers({ 'Content-Type': 'application/json' }),
        };

        return fetchUtils.fetchJson(url, options);
    },
    getMany: (resource, params) => {
        // make a POST request to the backend API
        const url = `${apiUrl}/${resource}`;
        return httpClient(url).then(({ json }) => ({
            data: json,
            total: json.length,
        }));
    },
    getManyReference: (resource, params, target) => {
        const url = `${apiUrl}/${resource}?server_id=${params.id}`;
        return httpClient(url).then(({ json }) => ({
            data: json,
            total: json.length,
        }));
    },
};

export default dataProvider;