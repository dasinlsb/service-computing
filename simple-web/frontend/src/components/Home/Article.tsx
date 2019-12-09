import React, {useContext, useEffect, useReducer} from 'react';
import CssBaseline from '@material-ui/core/CssBaseline';
import Typography from '@material-ui/core/Typography';
import { makeStyles } from '@material-ui/core/styles';
import {Paper} from "@material-ui/core";
import {Redirect} from "react-router";
import {useAuth} from "../AuthProvider";

import Grid from "@material-ui/core/Grid";
import Link from "@material-ui/core/Link";


const useStyles = makeStyles(theme => ({
  articlePaper: {
    padding: theme.spacing(3, 2),
    marginBottom: theme.spacing(3),
  },
}));

interface ArticleData {
  title: string;
  content: string;
  floors: number;
  url: string;
}

interface ArticleProps {
  location: any,
}

type ReducerAction =
  |{ type: 'goto_page';
     target: number;
  }
  |{ type: 'refresh_data';
     status: boolean;
     data: ArticleData[];
  }

interface ArticleState {
  hasFetched: boolean;
  articleList: ArticleData[];
  page: number;
}

const initState: ArticleState = {
  hasFetched: false,
  articleList: ['?', '?'].map(s => ({
    'title': s,
    'content': s,
    'floors': 0,
    'url': 'https://emacs-china.org',
  })),
  page: 1,
};

function reducer(state: ArticleState, action: ReducerAction) {
  switch (action.type) {
    case 'goto_page':
      return { hasFetched: false, articleList: [], page: action.target, };
    case 'refresh_data':
      return { ...state, hasFetched: action.status, articleList: action.data, };
    default:
      return state;
  }
}

const Article: React.FC<ArticleProps> = (props: any) => {
  const classes = useStyles();
  const [state, dispatch] = useReducer(reducer, initState);
  const auth = useAuth();
  const fetchData = () => {
    auth.fetchArticleList(state.page)
      .then(
        (articles: ArticleData[]) => {
          dispatch({ type: 'refresh_data', status: true, data: articles });
        }
      )
      .catch(e => {
        console.log('Fetch article data failed: ', e);
        dispatch({ type: 'refresh_data', status: true, data: state.articleList });
      })
  };

  useEffect(() => {
      fetchData();
  }, []);

  if (!auth.state.isAuthenticated) {
    return <Redirect to={'/login'} />
  }

  const article = state.articleList.map(({title, content, floors, url}: ArticleData, index: number) => (
    <Paper className={classes.articlePaper} key={index} onClick={() => console.log('go')}>
      <Grid container>
        <Grid item xs={8}>
          <Typography variant="h5" component="h3">
            <Link href={url}>
              {title}
            </Link>
          </Typography>
        </Grid>
        <Grid item xs={4}>
          <Typography component="p">
            {floors}
          </Typography>
        </Grid>
      </Grid>
    </Paper>
  ));
  return (
    <React.Fragment>
      <CssBaseline />
      {article}
    </React.Fragment>
  );
};

export default Article;
