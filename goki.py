import re
import itertools
from tqdm import tqdm
import numpy as np
import pandas as pd
import matplotlib.pyplot as plt
from hazm import *
from hazm.utils import stopwords_list as persian_stopwords
import nltk
from nltk.corpus import stopwords
import html, string

nltk.download('stopwords')
english_stopwords = set(stopwords.words('english'))
class Normalization:
    def __init__(self, text: str) -> None:
        self.text = text
        self.tokenizer = SentenceTokenizer()
        #tokenize_message = lambda x: [sentence for sentence in tokenizer.tokenize(x) if not sentence.startswith('IP:')]

    # clean messages
    def clean(self):
        if (not self.text) or self.text.isdigit():
            self.text = self.text.replace("\u200c", "‌‌")
            self.text = self.text.replace(u'\xa0', u' ')
            self.text = self.text.replace(u'\t', u' ')
            # remove html special character

            self.text = html.unescape(self.text)
            self.text = self.text.lower()
            return self.text
        else:
            return None 
    
    def filter_finglish(self):
        self.text = self.clean(self.text)
        pattern = r"[a-z0-9\s'!\"#$%&\'()*+,-.\/:;<=>?@\[\\\]\^_`{\|}~']+$"
        if self.text != None:
            if bool(re.match(pattern, self.text)):
                return None
            else:
                return self.text

    def tokenize_stats_ip(self):
        if not 
