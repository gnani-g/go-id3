from mutagen.id3 import ID3 as _ID3
from mutagen.id3 import TIT2
import re

def hello():
    print "hello"


class ID3(object):
    def __init__(self, f, v2_v=4):
        self.wrapped_class = _ID3(f, translate=False, v2_version=v2_v)

    def __getattr__(self,attr):
        orig_attr = self.wrapped_class.__getattribute__(attr)
        if callable(orig_attr):
            def hooked(*args, **kwargs):
                # self.pre()
                result = orig_attr(*args, **kwargs)
                # self.post()
                
                # prevent wrapped_class from becoming unwrapped
                if result == self.wrapped_class:
                    return self
                
                return result
            return hooked
        else:
            return orig_attr

    def pre(self):
        print ">> pre"

    def post(self):
        print "<< post"

    def save(self, v2_v=4):
        self.wrapped_class.save(v2_version=v2_v)

def test():
    i = ID3("dingdong.mp3", 3)
    told = i.getall("TIT2")[0]
    # print 
    
    # re_pattern = re.compile(u'[^\u0000-\uD7FF\uE000-\uFFFF]', re.UNICODE)
    # filtered_string = re_pattern.sub(u'\uFFFD', told.text)  

    t = TIT2(1, told.text[0].decode('ascii'))
    i.setall("TIT2", [t])
    i.save(3)
    ''' print i.getTitle() '''
