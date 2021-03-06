import os

import pandas as pd
from sklearn import datasets
from sklearn import model_selection


if __name__ == '__main__':

    X, y = datasets.load_breast_cancer(return_X_y=True)

    df = pd.DataFrame(X)
    df['has_cancer'] = y

    train, val = model_selection.train_test_split(df, test_size=0.33, random_state=10)

    here = os.path.dirname(os.path.realpath(__file__))
    train.to_csv(os.path.join(here, 'train.csv'), index=False)
    val.to_csv(os.path.join(here, 'val.csv'), index=False)
