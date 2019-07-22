import ReactDOM from 'react-dom';
import {App} from './components/App';
import * as React from "react";
import {CssBaseline} from "@material-ui/core";
import {ThemeProvider} from '@material-ui/styles';
import theme from "./theme";

ReactDOM.render(
    <ThemeProvider theme={theme}>
        {/* CssBaseline kickstart an elegant, consistent, and simple baseline to build upon. */}
        <CssBaseline/>
        <App/>
    </ThemeProvider>,
    document.getElementById('root')
);
