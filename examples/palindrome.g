def "is_palindrome" value
  (str: NonEmptyString)
    fn
      def "arr" value
        chars str end
      end
      
      def "l" value
        .length arr end
      end

      def "k" value 0 end

      def "s" value "" end
      def "e" value "" end

      def "is_p" value true end

      while k < l
        do
          let "s" value
            array.read arr k end
          end

          let "e" value
            array.read arr l - k - 1 end
          end

          when
            s != e
          then
            let "is_p" value false end
          end

          let "k" value k + 1 end
        end
      end

      is_p
    end
end

def "words" value
  array
    "racecar"
    "hannah"
    "kayak"
    "radar"
    "saippuakivikauppias"
    "tattarrattat"
    "achilles"
    "gecko"
    "golang"
    "language"
    "fun"
  end
end


array.for words
  (w i)
    fn
      std.write
        .is_palindrome w end
      end
    end
end
