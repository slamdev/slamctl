import React from 'react';
import Container from '@material-ui/core/Container';
import Typography from '@material-ui/core/Typography';
import Box from '@material-ui/core/Box';
import {ProTip} from "./ProTip";
import {Link} from "@material-ui/core";

const MadeWithLove: React.FunctionComponent<{}> = () =>
    <Typography variant="body2" color="textSecondary" align="center">
        {'Built with love by the '}
        <Link color="inherit" href="https://material-ui.com/">
            Material-UI
        </Link>
        {' team.'}
    </Typography>;

export const App: React.FunctionComponent<{}> = () =>
    <Container maxWidth="sm">
        <Box my={4}>
            <Typography variant="h4" component="h1" gutterBottom>
                Create React App v4-beta example with TypeScript
            </Typography>
            <ProTip/>
            <MadeWithLove/>
        </Box>
    </Container>;
